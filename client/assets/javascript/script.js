"use strict"

let title = document.getElementById('pageTitle').innerHTML
let homeLink = document.getElementById('home')
let registerLink = document.getElementById('register')
let loginLink = document.getElementById('login')
let logOutLink = document.getElementById('logout')


function managePageDisplay(title) {
	console.log(title)
	switch (title) {
		case 'register':
			registerLink.style.display = 'none';
			logOutLink.style.display = 'none';
			break;
		case 'login':
			loginLink.style.display = 'none';
			logOutLink.style.display = 'none'
			break;
		case 'profile':
			registerLink.style.display = 'none';
			break;
		case 'home':
			homeLink.style.display = 'none';
			break;
		default:
			break;
	}
}

managePageDisplay(title);