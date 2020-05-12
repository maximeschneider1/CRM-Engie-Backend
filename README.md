# Projet data - back-end

Description

This project is a CRM made for EngieMyPower service, which provides clients help dealing with their personal solar installations.
It contains all useful tool for the Engie employee that deals with customers to manage their needs and find their problem. The goal is to have a simple yet effective CRM which allows to manage clients and leads.

Features :
- Clients and leads management 
- Todo-list system 
- Alert system 
- Documentation system 
- Tag management system : customize push marketing for clients and leads
- Client scoring algorithm : determine where are the clients that needs help 
- Lead scoring algorithm : determine the most "hot" leads
- Employee KPI 
- Clients and leads KPI 

How to configure and run the code : 
This project needs to have a config.json file in the /config folder with access codes to the database. A model of what config.json file should look like is already present in /config at config-dist.json


Prerequisites

You need to pocess go version go1.13.5 darwin/amd64 to run this code

To run the code : 

go run main.go 


Functionnement by packages `

- config : responsible for reading the config and returning a db connection. 
- dao : the main source of data for the application.
- handler : package responsible for launching a web server and handling requests
- model : contain the model for the most important data structures used in the application
- service : contains the scoring algorithms and the predictive maintenance algorithms. 

Please note that a detailed functioning by function is available with comments for all exported functions

Bugs (front-end)
- Filtering through the articles 
- Filtering through a client solar production 

Bugs(back-end) 
- Returning the same clients and leads score in the list and on the detailled page

Next development : 
Within our project for Engie, we planned for them to do another version of the CRM focused on the visualisation of the strategic, macro aspects of Engie My Power.  






