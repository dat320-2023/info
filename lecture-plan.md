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
|35.36| 29.08| Tue | 6  | [Mechanism: Limited Direct Execution]                                                |
|    |       | Tue | 7  | [Scheduling: Introduction]                                                           |
|    | 30.08 | Wed |    | [Some Go coding] (demo)                                                              |
|    |       | Wed | 8  | [Scheduling: Multi-Level Feedback Queue]                                             |
|    |       | Wed | 9  | [Scheduling: Proportional Share]                                                     |
|    |       | Wed |    | [Scheduling: Test questions]                                                         |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
|37.38|    | Tue | 13 | [Abstraction: Address Spaces]                                                        |
|    |       | Tue | 14 | [Memory API]                                                                         |
|    |    | Wed | 15 | [Mechanism: Address Translation]                                                     |
|    |       | Wed | 16 | [Segmentation]                                                                       |
|    |       | Wed | 17 | [Free-Space Management]                                                              |
|    |       | Wed |    | [Test questions]                                                                     |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
|37.38|  | Tue | 18 | [Paging: Introduction]                                                               |
|    |       | Tue | 19 | [Paging: Faster Translation (TLBs)]                                                  |
|    |   | Wed | 20 | [Paging: Smaller Tables]                                                             |
|    |       | Wed | 21 | [Beyond Physical Memory: Mechanisms]                                                 |
|    |       | Wed |    | [Test questions]                                                                     |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 39 |   | Tue |    | [Recap and coding]                                                                   |
|    |       | Tue |    | [Recap and coding]                                                                   |
|    |   | Wed |    | [Recap and coding]                                                                   |
|    |       | Wed |    | [Questions]                                                                          |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 40 |  | Tue |    | [Recap / Exercise]                                                                              |
|    |       | Tue |    | [Recap / Exercise]                                                                            |
|    |  | Wed |    | [Recap / Exercise]                                                                             |
|    |       | Wed |    | [Recap / Exercise]                                                                       |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 41 |  | Tue | 26 | [Concurrency: Introduction]                                                          |
|    |       | Tue | 27 | [Thread API]                                                                         |
|    |  | Wed |    | [Rob Pike: Go Concurrency Patterns][1]                                               |
|    |       | Wed |    | [Live Coding: Shared Integer w/Mutual Exclusion]                                     |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 42 |  | Tue | 28 | [Locks]                                                                              |
|    |       | Tue | 29 | [Lock-based data structures]                                                         |
|    |  | Wed | 30 | [Condition Variables]                                                                |
|    |       | Wed | 31 | [Semaphores]                                                                         |
|    |       | Wed | 32 | [Common Concurrency Problems]                                                        |
|    |       | Wed | 33 | [Event-based Concurrency]                                                            |
|    |  | Wed |    |                                                                                      |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
|  43  |  |  |    | [No Lecture]                                                                 |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 44 |  | Tue | 10 | Multiprocessor Scheduling                                                            |
|    |       | Tue |    | [Concurrency in go]                                                                  |
|    |  | Wed |    | [Concurrency in go ]                                                                 |
|    |  | Wed |    |                                                                                      |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 45 |  | Tue | 36 | [I/O Devices]                                                                        |
|    |       | Tue |    |                                                                                      |
|    |  | Wed | 37 | [Hard Disk Drives]                                                                   |
|    |       | Wed |    |                                                                                      |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 46 |  | Tue |    | [Recap]                                                                              |
|    |       | Tue |    | [Recap]                                                                              |
|    |  | Wed |    | [Recap]                                                                              |
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
