#!/bin/sh

cd ./frontend/;
npm run dev & $(cd ../backend/; go run main.go);
