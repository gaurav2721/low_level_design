Following are the steps for an LLD 
1. Gather functional and non functional requirements
2. Write down the functionality in terms of functions and objects 
3. write down your code in form of objects(OOPS and SOLID , KISS(keep it simple stupid) , YAGNI(You aint gonna need it ))
4. There must be an api interface to your module , use api interface and packages and functions to split your code into manageable chunks 
Game Engine

Functional Requirement
1. Taking a game and having the ability to play it through AI  

Api Design
1. func Initialize(game string)(Board,err)
2. GameObject.RegisterPlayers(player1 Player,player2 Player)(error)
3. Player1.Move(gameObject GameObject,move Move)(bool,error)

While defining our api interfaces we will write some objects and we will have to define them later on 

When you are defining functions ask yourself this question , what are the object and what jobs will they be doing that in the starting start with just simple
functions
splitting the code into functions to reduce redundancy and have lightly coupled code 
create separate files for different classes and rearranging them in packages 
Think about what are the apis that you are going to expose --> GameEngine
Create packages 
