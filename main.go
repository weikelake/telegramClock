package main

import (
	"context"
	"fmt"
	"github.com/anonyindian/gotgproto"
	"github.com/anonyindian/gotgproto/dispatcher"
	"github.com/anonyindian/gotgproto/sessionMaker"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
	"log"
	"telegramClock/clock"
	"telegramClock/settings"
	"time"
)

func main() {
	dp := dispatcher.MakeDispatcher()
	gotgproto.StartClient(gotgproto.ClientHelper{
		AppID:      settings.GetUserData().AppId,
		ApiHash:    settings.GetUserData().ApiHash,
		Session:    sessionMaker.NewSession("session_name", sessionMaker.Session),
		Phone:      settings.GetUserData().Phone,
		Dispatcher: dp,
		TaskFunc: func(ctx context.Context, client *telegram.Client) error {
			go func() {
				for {
					if gotgproto.Sender != nil {
						fmt.Println("client has been started")
						clock.GenerateClockPicture()
						var profilePhoto tg.PhotosPhoto
						for {
							if !isRoundedTime() {
								time.Sleep(1 * time.Second)
								continue
							}

							var (
								profilePhotoForDelete tg.PhotosPhoto
								err                   error
							)
							profilePhotoForDelete = profilePhoto

							profilePhoto, err = uploadProfilePhoto(ctx, client, settings.GetPicturePath())
							if err != nil {
								fmt.Println(err)
								log.Fatalln(err)
								break
							}

							time.Sleep(30 * time.Second)

							go clock.GenerateClockPicture()

							if profilePhotoForDelete.Photo == nil {
								continue
							}
							err = deleteProfilePhoto(ctx, client, profilePhotoForDelete)

							if err != nil {
								fmt.Println(err)
								log.Fatalln(err)
								break
							}

						}
						break
					}
				}

			}()
			return nil
		},
	})
}

func isRoundedTime() bool {
	if time.Now().Second() == 0 {
		fmt.Println("rounded")
		return true
	}
	return false
}

func deleteProfilePhoto(ctx context.Context, client *telegram.Client, profilePhoto tg.PhotosPhoto) error {
	fmt.Println("delete")
	raw := client.API()
	photoStruct, _ := profilePhoto.GetPhotoAsNotEmpty()

	input := photoStruct.AsInput()

	var ipc tg.InputPhotoClass
	ipc = input

	var photosArray []tg.InputPhotoClass
	photosArray = append(photosArray, ipc)

	_, err := raw.PhotosDeletePhotos(ctx, photosArray)
	if err != nil {
		return err
	}
	return nil
}

func uploadProfilePhoto(ctx context.Context, client *telegram.Client, filename string) (tg.PhotosPhoto, error) {
	fmt.Println("upload")
	raw := client.API()
	u := uploader.NewUploader(raw)
	file, err := u.FromPath(ctx, filename)
	if err != nil {
		return tg.PhotosPhoto{}, err
	}

	profilePhoto, err := raw.PhotosUploadProfilePhoto(ctx, &tg.PhotosUploadProfilePhotoRequest{
		File: file,
	})

	if err != nil {
		return tg.PhotosPhoto{}, err
	}

	return *profilePhoto, nil
}
