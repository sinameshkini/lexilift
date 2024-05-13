# LexiLift

LexiLift is a free and open-source CLI app designed to help you learn any English word you want!

```text
$ ./lexilift                                                                                                      
  _                     _   _       _    __   _
 | |       ___  __  __ (_) | |     (_)  / _| | |_
 | |      / _ \ \ \/ / | | | |     | | | |_  | __|
 | |___  |  __/  >  <  | | | |___  | | |  _| | |_
 |_____|  \___| /_/\_\ |_| |_____| |_| |_|    \__|

LexiLift is a free and open-source CLI app designed to help you learn any English word you want!

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> LexiLift <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Menu:
	1- Review my words
	2- Add a new word to my words
	3- Add words list to my words
	4- Review history
	m- Menu
	d- Dashboard
	c- Clear
	q- close the app
Press the character corresponding to the action you want to perform
:::> d
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Dashboard <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Most Score Words:
	1- evidence	    (Score: 12, Proficiency: 6)
	2- skyscraper	(Score: 12, Proficiency: 6)
	3- decade	    (Score: 12, Proficiency: 6)
	4- consume	    (Score: 11, Proficiency: 6)
	5- consumption	(Score: 10, Proficiency: 6)

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
    24- 2024-05-12 00:57  26s	FP:-4	TP:-3	CNT:6	KNW:6	NK:0	S:39
    23- 2024-05-12 00:45  16s	FP:3	TP:3	CNT:2	KNW:2	NK:0	S:10
    22- 2024-05-12 00:16  1m12s	FP:5	TP:5	CNT:13	KNW:9	NK:1	S:91
    21- 2024-05-11 22:19  5m58s	FP:-7	TP:6	CNT:100	KNW:63	NK:34	S:25
    20- 2024-05-11 21:14  9m17s	FP:-7	TP:6	CNT:100	KNW:63	NK:34	S:15
    Total: 24, Duration: 1h55m27s, Score: 170
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
- [x] Calculate score in reviews
- [x] Track review count for each word
- [x] Add option to add a new word during review
- [x] Show user's words sorted by proficiency
- [x] Display total review duration in the dashboard
- [x] review comment
- [ ] manage reviews
