liblibrary.so: dllmain.cpp
	gcc -o libmain.dll dllmain.cpp -std=c++17 -O3 -Wall -Wextra -fPIC -shared

fancy: main.go
	go build -x -buildmode=c-shared -o main-fix.dll main.go