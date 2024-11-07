# PESU IO - Go Final Project - Making the BackEnd of Leetcode
The basis of this project is to make a backend, which can support the managing and storage of coding questions based with the help of gin, gorm and go; to mimic the backend of Leetcode.
The primary function of this project is to POST questions, which will have : 
1. TestCodes - To compare the actual answer with the output from the code written by the user; *in certain cases, multiple inputs may be provided so it may compare both input and output to check the corresponding solution*
2. Score - How much the answer matches with the output (on a scale from 1 - 10)
3. Difficulty - How hard a question is, categorized by the user (easy, medium & hard)


# Pre-Requisites
1. [GOLang](https://go.dev/)
2. [Gin-Gonic](https://github.com/gin-gonic/gin)
3. [GORM](https://gorm.io/)

# API Endpoints

| Endpoint         | Method | Path             | Request Example                                                                                                         |
|------------------|--------|------------------|-------------------------------------------------------------------------------------------------------------------------|
| *Sign Up*      | POST   | /auth/signup   | json {"username": "testuser", "password": "password123"}                                                         |
| *Sign In*      | POST   | /auth/signin   | json {"username": "testuser", "password": "password123"}                                                         |
| *Create Question* | POST | /question/create | json {"question": "What is the capital of France?", difficulty : "easy", "testCases": [{"input": "", "expectedOutput": "Paris"}], "score": 10} |
| *Fetch Question* | POST | /question/fetch | json {"questionID": 1}                                                                                           |
| *Run Code (Python)* | POST | /run           | json {"language": "python", "code": "print('Hello World')", "input": ""}                                         |
