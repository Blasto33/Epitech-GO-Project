## Epitech Go Project

A Go project made for Epitech.

### How to compile
`Warning you need to specify in the dockerfile the name of the file that you wanna use. By default it's file.txt`

`docker build . -t warehouse-project`

`docker scan warehouse-project`
 
### How to run the project

`docker run warehouse-project`

### Architecture

We got a package for the parsing and a package for the algorithm.
The algorithm will call the package which parses to execute it and get the infos it needs.

### Strategy

For each packet, we are looking for the closest palette. The palette will move to the packet but if she finds another packet on her way, she will take it instead of the original one.
Then each palette with a packet is looking for the closest truck. They will wait near the truck if he's gone.
The truck leave only if he can't take a blue packet anymore (500) or if there is no packet left.

### Component Diagram

![C4 Diagram](https://i.imgur.com/Ds5ghU9.png)
