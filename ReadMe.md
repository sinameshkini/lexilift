# LexiLift
Is a free and open source CLI app for learning any English word you want!

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

## Install
visit release page: https://github.com/sinameshkini/lexilift/releases
download executable binary file and run it in your terminal.

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

## TODO
- [ ] complete docs
- [ ] test and finalize audio
- [ ] build with tags
- [ ] build for all golang supported platforms
- [ ] unit test
- [ ] integration test
- [ ] multi language
- [ ] custom configuration (default lang, word view, etc)
- [ ] word tags
- [ ] my words management
- [ ] review by tags
- [ ] show more info like crated at, pronunciation, synonyms and antonym of word in review (existed in db)
- [ ] impl Data completion words operation
- [ ] import/export words
- [ ] calculate points in review (0 to 100)
- [ ] count review each word
- [x] add new word option in review running
- [x] show my words sorted by proficiency
- [x] show total review duration in dashboard
- 