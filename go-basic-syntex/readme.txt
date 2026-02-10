# exercise
https://www.w3schools.com/go/exercise.php

# สร้าง module ให้โปรเจค module เป็นที่รวบรวม package ที่ใช้ใน project

syntax: 
go mod init <package name>
ex:
go mod init pwa-workshop/go-basic-syntax

# import package ภายนอก
syntax: 
go get <package name>
ex:
go get github.com/google/uuid


# รันโปรแกรมแบบไม่ได้ไฟล์ binary
syntax: 
go run <file name>
ex: 
go run main.go
go run .  # ใช้ได้ถ้ามีการสร้าง module ให้ project แล้ว

# รันโปรแกรมแบบได้ไฟล์ binary
syntax:
go build -o <app name>
ex:
go build -o app



# golang structure
myapp/                ← Go module
├─ go.mod
├─ go.sum
├─ main.go            ← package main
├─ handler/           ← package handler
│  └─ user.go
├─ service/           ← package service
│  └─ user.go
└─ repository/        ← package repository
   └─ user.go
