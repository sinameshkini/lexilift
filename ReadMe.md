# LexiLift

LexiLift is a free and open-source CLI app designed to help you learn any English word you want!

```text
$ ./lexilift                                                                                                      
  _                     _   _       _    __   _
 | |       ___  __  __ (_) | |     (_)  / _| | |_
 | |      / _ \ \ \/ / | | | |     | | | |_  | __|
 | |___  |  __/  >  <  | | | |___  | | |  _| | |_
 |_____|  \___| /_/\_\ |_| |_____| |_| |_|    \__|

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Dashboard <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
My Words:
Proficiency: -6	 Count: 5
Proficiency: -5	 Count: 1
Proficiency: -4	 Count: 8
Proficiency: -2	 Count: 7
Proficiency: -1	 Count: 40
Proficiency: 0	 Count: 1
Proficiency: 1	 Count: 15
Proficiency: 2	 Count: 2
Proficiency: 3	 Count: 19
Total: 98

My Reviews:
Total: 12, Duration: 1h20m10s
12- 2024-05-10 20:23  9m17s	FP:0	TP:3	CNT:61	KNW:33	NK:25
11- 2024-05-10 20:15  8m28s	FP:-5	TP:-1	CNT:37	KNW:21	NK:16
10- 2024-05-10 15:27  5m21s	FP:-6	TP:-2	CNT:30	KNW:14	NK:16

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> LexiLift <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Menu:
	0- Dashboard
	1- Review my words
	2- Add a new word to my words
	3- Add words list to my words
	4- Review history
	m- Menu
	c- Clear
	q- close the app
Press the character corresponding to the action you want to perform
:::> 
```

## Installation

To install LexiLift, visit the release page: [LexiLift Releases](https://github.com/sinameshkini/lexilift/releases), download the executable binary file, and run it in your terminal.

## Installation for Developers

### Install Dependencies
```shell
sudo apt install git make libasound2-dev
```

### Install Golang (version 1.22.0 or higher)
Install Golang on your workstation from [the official website](https://go.dev/dl/).

### Clone Repository
```shell
git clone https://github.com/sinameshkini/lexilift
```

### Run
```shell
cd lexilift

# Method 1: Build and run
make build
./lexilift

# Method 2: Run directly
make run
```

## To-Do List
- [ ] Complete documentation
- [ ] Test and finalize audio functionality
- [ ] Build with tags
- [ ] Build for all Golang-supported platforms
- [ ] Implement unit tests
- [ ] Implement integration tests
- [ ] Add multi-language support
- [ ] Implement custom configuration (default language, word view, etc.)
- [ ] Add word tags functionality
- [ ] Implement management for user's words
- [ ] Allow reviewing words by tags
- [x] Display additional information such as creation date, pronunciation, synonyms, and antonyms of words in reviews
- [ ] Implement data completion for words operation
- [ ] Add import/export functionality for words
- [ ] Calculate points in reviews (0 to 100 scale)
- [x] Track review count for each word
- [x] Add option to add a new word during review
- [x] Show user's words sorted by proficiency
- [x] Display total review duration in the dashboard

---