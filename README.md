Following are the steps for an LLD [Readable , maintainable and extensible code / Behvaiour driven tests ]
1. Gather functional and non functional requirements
2. Write down the functionality in terms of functions and objects 
4. There must be an api interface to your module , use api interface and packages and functions to split your code into manageable chunks 
5. The relationship between objects can be one to one , one to many and many to many , or object may be independent , there may be parent child relationship 
7. try not to have cyclic dependency amongst packages 
11. Public apis in each object [this should have input validation as well ]

13. concurrent access
17. Handling packages post doing the code changes makes sense 

OOPS
1. Abstration - hiding member variable from outside entities
2. 

SOLID
Single responsibility - A function does one and one task only , a class has only one task [review if it should have all the information inside it which we want it to have ,it does not take extra functionality inside it ]
Open Closed Principle - you have a core logic which is closed for modification , but you can add more functionalities to it , using interfaces 
// Liskov substitution principle :- if class b is subtype of class A , then we should be able to replace object of class A with class B without breaking the program
// Bike interface with two objects bike and cycle , but turnOnEngine will not be present in cycle , however increaseSpeed will be present in both
// To fix for Liskov Substitution Principle : we split classes for eg Vehicle is parent and EngineVehicle and Bicycle are subclass and EngineVehicle has subclasses Car and 

// Interface Segregation Principle :- A client should not be forced to depend on interfaces they do not need  for example
// Restaurant Employee (take order , make dish ) -> split into two interface waiter interface and chef interface

// Dependency inversion principle :- a class should depend on interfaces , rather that concrete classes , for eg if there is a macbook class , it should use iMouse and iKeyboard,
// for eg these can be mouse can be a wired and bluetooth one

SO --> Are always strived for

LID --> Related to inheritance 

KISS
-> We must try to accept generic class/interface and return subClass

DRY - method inlining , you have to actively think what piece of code can be common and then use that 
Extraction
1. Using function pointers 
2. Type casting 
3. splitting the code into functions to reduce redundancy and have lightly coupled code 
4. Composition

YAGNI
- if you dont need a parameter in a function dont pass it 


Readability 
1. Try catch comment logs and constants and dont use an antiPattern(that is wrong name ) , Input Validation must be there 

Orchestrator , State manager , rules engine , config pattern and object pools 

Api and packages - rest Apis and no cyclic dependency in packages 



Design Patterns
1. Prototype design pattern - ab object being able to create a new object with state changes , make sure about deep and shallow copy
2. Iterator design pattern - may have to use function literals / lambda function - HasNext , Next ---> while doing this we need to make sure the code does not become tightly coupled for that we use open close principle 
3. Wrapper design patterns 
4. Builder design pattern 
5. Singleton design pattern - using interface object to create a code snippet for creating a singleton class 
6. The Chain of Responsibility is a behavioral design pattern used to pass a request along a chain of handlers. Each handler decides either to process the request or to pass it to the next handler in the chain.
7. Memmento Design pattern allows you to create history of object , you should implemet save , restore and edit functionalities
8. Proxy design pattern : In order to maintain history we keep the Object Proxy stored which is of less size than the every object 
9. Flyweight design pattern : sharing an object accross multiple objects to save on memory
10. Strategy design pattern : depending on the input you decide on the strategy to execute , there is usually no switching tof strategy
11. State Design pattern : It stores the state , action and those actions take you to the next state , that is it has action() and next() functions , there is a very clear DAG
12. Factory design pattern 
13. Command design pattern : We create and pass an object that encapsulates whatever we want to do 



Game Engine

Functional Requirement
1. Taking a game and having the ability to play it through AI  
2. Game can be a board only as of now 

Thought process 
1. First write down the functions and objects 
2. api interface[one of the ways to think about is GET,POST,PUT,PATCH] , data models and package
3. relationship b/w objects
4. try catch comment logs 
5. rules engine. state manager(think about if you want things to be stateless or not ) and orchestrator 
6. input validation for a function

Things not to do
1. dont have cyclic dependency
2. 

Api Design
1. func Initialize(game string)(Board,err)
2. GameObject.RegisterPlayers(player1 Player,player2 Player)(error)
3. Player1.Move(gameObject GameObject,move Move)(bool,error)

While defining our api interfaces we will write some objects and we will have to define them later on 

When you are defining functions ask yourself this question , what are the object and what jobs will they be doing that in the starting start with just simple
functions

create separate files for different classes and rearranging them in packages 
Think about what are the apis that you are going to expose --> GameEngine
Create packages 

Marks are assigned on the basis of 
1. We are optimizing for readability , extensibility and correctness , maintainable , avoid tight coupling , backward compatibility
