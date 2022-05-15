# control-direction-withGo
Hello every one! This project is to control a rover that explore Mars planet.the mission control
already gave us a rover instructions to move into specific directions and positions.
# Rules
Step1: Write a code that read a file in specific format.<br />
&nbsp;&nbsp;&nbsp;&nbsp;a. First line will always be a size of a maps represent in integer only.<br />
&nbsp;&nbsp;&nbsp;&nbsp;b. After that will be an instruction to move or rotate.<br />
&ensp;&ensp;&ensp;&ensp;i. F for moving forward 1 block.<br />
&ensp;&ensp;&ensp;&ensp;ii. L for turning left.<br />
&ensp;&ensp;&ensp;&ensp;iii. R for turing right.<br />
&ensp;&ensp;&ensp;&ensp;iv. After the turning it should stay on the same block. Just direction is move.<br />
Step2: Rover always start at 0,0 with facing north.<br />
Step3: Rover cannot be in negative position.<br />
Step4: Rover should always report a current direction and position in format of {direction}:
{position-x},{position-y}<br />
&nbsp;&nbsp;&nbsp;&nbsp;a. N for North<br />
&nbsp;&nbsp;&nbsp;&nbsp;b. E for East<br />
&nbsp;&nbsp;&nbsp;&nbsp;c. W for West<br />
&nbsp;&nbsp;&nbsp;&nbsp;d. S for South<br />
Step5: If rover reaching the edge it must maintain the direction and position.<br />
# Example Input/Output
![image](https://user-images.githubusercontent.com/105012524/168476762-cbf808b8-feeb-4bf3-a185-9430495e7bdd.png)<br />

Here is the reference for read text file in Golang<br />
https://gosamples.dev/read-file/#:~:text=The%20simplest%20way%20of%20reading,by%20line%20or%20in%20chunks.
