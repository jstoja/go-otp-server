var timer = 0, tt;
function Init() {
	var date = new Date;
	tt = setInterval(Clock, 1000);
	if (date.getSeconds() === 30 || date.getSeconds() == 0) {
		Clock();
	} else {
		var savedTime = date.getSeconds();
		if (savedTime > 30) {
			savedTime -= 30;
		}
		Clock();
		timer = 30-savedTime;
	}
}
function SendASP() {
	xmlhttp = new XMLHttpRequest();
	xmlhttp.open("GET", "/otp", false);
	xmlhttp.send(null);
	return (eval(xmlhttp.responseText));
}
function ActualizeKeys() {
	tokens = SendASP();
	console.log(tokens);
	var keys = document.getElementById('keys').children;
	for (var i = keys.length - 1; i >= 0; i--) {
		keys[i].children[1].innerText = tokens[i];
	};
}
function Clock() {
	if (timer == 0) {
		ActualizeKeys();
		timer = 30;
	}
	console.log(timer);
	timer--;
}
Init();