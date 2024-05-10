# LexiLift
Is an free and open source CLI app for learning any English word you want!

```shell
$ ./lexilift                                                                                                      
  _                     _   _       _    __   _
 | |       ___  __  __ (_) | |     (_)  / _| | |_
 | |      / _ \ \ \/ / | | | |     | | | |_  | __|
 | |___  |  __/  >  <  | | | |___  | | |  _| | |_
 |_____|  \___| /_/\_\ |_| |_____| |_| |_|    \__|

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Dashboard <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
My Words:
Proficiency: -1	 Count: 22
Proficiency: 1	 Count: 23
Proficiency: -4	 Count: 1
Proficiency: 0	 Count: 2
Proficiency: -2	 Count: 13
Total: 61

My Reviews:
Total: 3
3- 2024-05-10 11:09  5m21s	FP:-3	TP:-1	CNT:36	KNW:8	NK:14
2- 2024-05-10 03:27  23s	FP:-2	TP:-2	CNT:3	KNW:1	NK:1
1- 2024-05-10 03:20  23s	FP:-1	TP:-1	CNT:34	KNW:2	NK:1

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> LexiLift <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Menu:
	0- Dashboard
	1- Review my words
	2- Add a new word to my words
	3- Add words list to my words
	4- Review history
	q- close the app
press number of action you want do:

```

## Install for Developers

### Install dependency
```shell
sudo apt install git make
```

### Install Golang (go1.22.0 or higher)
install golang on your workstation with: https://go.dev/dl/

### Clone repository
```shell
git clone https://github.com/sinameshkini/lexilift
```

### Run
```shell
cd lexilift

# method 1: build and run
make build
./lexilift

# method 2: run directly
make run
```