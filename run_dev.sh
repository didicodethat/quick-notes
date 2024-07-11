#!/bin/sh

cd ./next/;
npm run dev & $(cd ../backend/; go run main.go);
