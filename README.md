# Event reminder

Do not forget important events!<br />
Write it down once and **NEVER** forget it.

<img align='center' alt = 'Automation' width = '300' src = 'https://assets.website-files.com/5daaade3e3e3f04da71daa8e/62905b592241713298ec0337_Event%20reminder%20Compress.gif'>

## About

The application is divided into frontend and backend part. In order for the application to work, three servers are started:

- the server on the [frontend](/front/)
- the server on the [backend](/server/)
- the server for the mysql database

For each user, his credentials and the events he creates are remembered in database. In order to make sure that the user still uses his account after logging in and not someone else, sessions are used, i.e. access and refresh tokens.

## Frontend

The frontend is made with the **React js**. <br />
All code related to the frontend is located in the [front](/front/) folder.

## Backend

The backend is made with the programming language **Go**.<br />
All code related to the backend is located in the [server](/server/) folder. 
Communication with the database is done through the backend.

## Login and Register

The first page that is launched when starting the application is the Login page. If the user already has an account, he needs to enter his credentials. There is a verification system that checks whether the credentials are entered correctly, or whether such an account already exists in the system. After successful login, the user is sent to the Home screen.

<p align="left">
  <img src="./images/loginPage.png" width="30%" align="center"/>
  <img src="./images/wrongLoginPage.png" width="28.5%" align="center"/>
</p>

If the user does not have an account, he needs to select the Sign Up option at the bottom of the screen, which will take him to the registration screen. During registration, it is necessary to fill in each field so that it meets the requirements defined for it, otherwise the registration will not be successful. Also, it is necessary that the user with that username does not exist in the database in order to be entered into the system.

<p align="left">
  <img src="./images/registerPage.png" width="30%" align="center"/>
  <img src="./images/wrongRegisterPage.png" width="27.4%" align="center"/>
</p>

If the registration is successful, a specific message with the Sign in option is displayed on the screen to transfer the user to the Login page.

<p align="left">
  <img src="./images/successPage.png" width="30%" align="center"/>
</p>

## Home

Since the user has successfully logged in, the Home screen is in front of him. It is made up of 3 parts:</br>
- The first part represents a **timer** that ticks and shows how much time is left until the first next event that occurs.
- The second part presents a **list of events** created by the user. The events are sorted by time so that the most recent events are displayed first. With the help of blue arrows, it is possible to view other events that the user has.
- The third part presents a **calendar** with which the user can find his way more easily in order to know when he wants to make an event.
  
<p align="left">
  <img src="./images/homePage.png" width="70%" align="center"/>
</p>

When the timer ticks, a notification will be displayed that it is time for the event and the timer will start counting down the time for the first next event:

<p align="left">
  <img src="./images/eventFinished.png" width="70%" align="center"/>
</p>

## Deleting an event from the list

Each event in the list has a sticker bin in the corner. By clicking on it, a pop-up window appears where the user is once again asked if he is sure he wants to delete the event.

<p align="left">
  <img src="./images/deleteEvent.png" width="70%" align="center"/>
</p>

## Adding event to the list

By clicking on the purple circle on which the pencil is drawn, the option to create a new event is launched. The user is able to choose the title and at what time he wants to schedule the event.

<p align="left">
  <img src="./images/newEventEmpty.png" width="30%" align="center"/>
  <img src="./images/newEventPicker.png" width="23.3%" align="center"/>
  <img src="./images/newEventPickerTime.png" width="22.8%" align="center"/>
</p>

