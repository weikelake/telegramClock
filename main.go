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
						for {
							clock.GenerateClockPicture()
							profilePhoto, err := uploadProfilePhoto(ctx, client, "clock/out.png")
							if err != nil {
								fmt.Println(err)
								break
							}
							time.Sleep(60 * time.Second)
							err = deleteProfilePhoto(ctx, client, profilePhoto)
							if err != nil {
								fmt.Println(err)
								break
							}
							time.Sleep(10 * time.Millisecond)
						}

						break
					}
				}
			}()
			return nil
		},
	})
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
