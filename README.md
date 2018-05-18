# ArrayProductSolution-AWSLambda
This is a solution to the problem "Given an array of numbers int[n] replace each number in the array with the product of all the other numbers in the array except the number itself WITHOUT using division. Example: [1,2,3,4] -> [24,12,8,6]"

I've also deployed this code as an aws lambda
--------------------------
Execution instructions
--------------------------
    POST https://2z6ck5bqaf.execute-api.us-east-1.amazonaws.com/dev/array
    BODY
    {
      "nums":[1,2,3,4]
    }
    
    RESPONSE
    {
      "result": "[24,12,8,6]"
    }
