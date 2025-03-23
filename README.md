# Go-Mobile-Applications
Tests gomobile 

 ../../gopath/bin/gomobile build -target=android -androidapi 26

gomobile build -androidapi=21

export PATH=$PATH:~/go/bin

[Gomobile Documentation](https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile)

[Gomobile Mirror](https://github.com/golang/mobile)

[Go bindings for OpenGL ES 2.0 and ES 3.0](https://pkg.go.dev/golang.org/x/mobile/gl)

[Mobile Wiki](https://go.dev/wiki/Mobile)

[fyne](https://fyne.io/)

 go install fyne.io/fyne/v2/cmd/fyne@latest

In Android Studio Go SDK Manager In the SDK Tool Install NDK Side by Side

Windows:
Edit System Environmental Variables and add

ANDROID_HOME
c:\Users\infor\appdata\local\android\sdk

ANDROID_NDK_HOME
  C:\Users\infor\appdata\local\android\sdk\ndk\29.0.13113456

echo %ANDROID_HOME%
echo %ANDROID_NDK_HOME%

Mac:
export [variable_name]=[variable_value]

../../gopath/bin/bin/fyne package -os android -appID com.example.myapp

/users/infor/go//bin/fyne package -os android -appID com.example.myapp
