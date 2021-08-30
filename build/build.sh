#sc create wsprefix binPath= "C:\Users\bbeltre\Documents\VS Code Projects\Go\wsprefix\wsprefix.exe"

#build1
go build

#build2
go build -ldflags "-s -w"

#build3
go build -ldflags "-s -w" -trimpath

#Create service
nssm.exe install  nameservice