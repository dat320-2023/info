# Lecture Plan 2023

## Book 
The course is based on this [textbook](http://pages.cs.wisc.edu/~remzi/OSTEP/):
Operating Systems: Three Easy Pieces. Version 1.00.
Remzi H. Arpaci-Dusseau and Andrea C. Arpaci-Dusseau

## Plan

This is a tentative lecture plan, it may change during the semester.

| W  | Date  | Day | Ch | Topics                                                                               |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 34 | 22.08 | Tue |    | [Course Intro]                                                                       |
|    |       | Tue |    | [Lab Intro] (demo)                                                                   |
|    | 23.08 | Wed | 2  | [Introduction to Operating Systems]                                                  |
|    |       | Wed | 4  | [Abstraction: The Process]                                                           |
|    |       | Wed | 5  | [Process API]                                                                        |
|    |       | Wed |    | [Some C coding] (demo)                                                               |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
|35.36| 29.08 | Tue | 6  | [Mechanism: Limited Direct Execution]                                               |
|    |       | Tue | 7  | [Scheduling: Introduction]                                                           |
|    | 30.09 | Wed |    | [Some Go coding] (demo)                                                              |
|    |       | Wed | 8  | [Scheduling: Multi-Level Feedback Queue]                                             |
|    |       | Wed | 9  | [Scheduling: Proportional Share]                                                     |
|    |       | Wed |    | [Scheduling: Test questions]                                                         |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 37 | 05.09 | Tue | 13 | [Abstraction: Address Spaces]                                                        |
|    |       | Tue | 14 | [Memory API]                                                                         |
|    | 06.09 | Wed | 15 | [Mechanism: Address Translation]                                                     |
|    |       | Wed | 16 | [Segmentation]                                                                       |
|    |       | Wed | 17 | [Free-Space Management]                                                              |
|    |       | Wed |    | [Test questions]                                                                     |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 38 | 12.09 | Tue | 18 | [Paging: Introduction]                                                               |
|    |       | Tue | 19 | [Paging: Faster Translation (TLBs)]                                                  |
|    | 13.09 | Wed | 20 | [Paging: Smaller Tables]                                                             |
|    |       | Wed | 21 | [Beyond Physical Memory: Mechanisms]                                                 |
|    |       | Wed |    | [Test questions]                                                                     |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 39 | 26.09 | Tue | 26 | [Concurrency: Introduction]                                                          |
|    |       | Tue | 27 | [Thread API]                                                                         |
|    | 27.09 | Wed |    | [Rob Pike: Go Concurrency Patterns][1]                                               |
|    |       | Wed |    | [Live Coding: Shared Integer w/Mutual Exclusion]                                     |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 40 | 03.10 | Tue | 28 | [Locks]                                                                              |
|    |       | Tue | 29 | [Lock-based data structures]                                                         |
|    | 04.10 | Wed | 30 | [Condition Variables]                                                                |
|    |       | Wed | 31 | [Semaphores]                                                                         |
|    |       | Wed | 32 | [Common Concurrency Problems]                                                        |
|    |       | Wed | 33 | [Event-based Concurrency]                                                            |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 41 | 19.09 | Tue |    | [Recap]                                                                              |
|    |       | Tue |    | [Recap]                                                                              |
|    | 20.09 | Wed |    | [Recap]                                                                              |
|    |       | Wed |    | [Questions]                                                                          |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 42 | 10.10 | Tue | 10 |  Multiprocessor Scheduling                                                             |
|    |       | Tue |    | [Concurrency in go]                                                                  |
|    | 11.10 | Wed |    | [Concurrency in go ]                                                                 |
|    | 11.10 | Wed |    |                                                                 |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 43 | 17.10 | Tue | 36 | [I/O Devices]                                                                        |
|    |       | Tue |    |                                                                                      |
|    | 18.10 | Wed | 37 | [Hard Disk Drives]                                                                   |
|    |       | Wed |    |                                                                                      |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 44 | 19.09 | Tue |    | [Recap]                                                                              |
|    |       | Tue |    | [Recap]                                                                              |
|    | 20.09 | Wed |    | [Recap]                                                                              |
|    |       | Wed |    | [Questions]                                                                          |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 45 | 31.10 | Tue |    | [Recap]                                                                              |
|    |       | Tue |    | [Recap]                                                                              |
|    | 01.11 | Wed |    | [Recap]                                                                              |
|    |       | Wed |    | [Questions]                                                                          |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 46 | 07.11 | Tue |    | [Recap]                                                                              |
|    |       | Tue |    | [Recap]                                                                              |
|    | 08.11 | Wed |    | [Recap]                                                                              |
|    |       | Wed |    | [Questions]                                                                          |

The following chapters are self study topics:
38,39,40

[1]: https://youtu.be/f6kdp27TYZs

## Syllabus also include

- published slides and code snippets available on [info](https://github.com/dat320-2023/info/)
- various videos and corresponding slides (see below) and
- lab assignments 1-7.

## Videos

- [Buffer Overflow (17:30)](https://youtu.be/1S0aBV-Waeo)
- [Spectre Attack Explained (3:41)](https://youtu.be/q3-xCvzBjGs)
- [Go Concurrency Patterns (51:26)](https://youtu.be/f6kdp27TYZs) ([slides](https://talks.golang.org/2012/concurrency.slide#1))
- More videos may be added.
