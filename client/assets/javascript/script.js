"use strict"

let title = document.getElementById('pageTitle').innerHTML
let homeLink = document.getElementById('home')
let registerLink = document.getElementById('register')
let loginLink = document.getElementById('login')
let profileLink = document.getElementById('profile')


function managePageDisplay(title) {
	console.log(title)
	switch (title) {
		case 'register':
			registerLink.style.display = 'none';
			break;
		case 'login':
			loginLink.style.display = 'none';
			break;
		case 'profile':
			profileLink.style.display = 'none';
			break;
		case 'home':
			homeLink.style.display = 'none';
			break;
		default:
			break;
	}
}

managePageDisplay(title);