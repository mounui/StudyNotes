function User(name) {
	this.name = name;
	this.isAdmin = false;
}

function Plant(name) {
	if (!new.target) {
		return new Plant(name);
	}
	this.name = name;
}

function Accumulator(startingValue) {
	this.startingValue = startingValue;
	this.read = function() {
		this.value = this.startingValue + +prompt('value?',0);
	}
}