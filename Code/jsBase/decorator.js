function cachingDecorator(func) {
	let cache = new Map();

	return function(x) {
		if (cache.has(x)) {
			return cache.get(x);
		}

		let result = func(x);

		cache.set(x, result);
		return result;
	};
}

function Calculator() {
	this.read = function() {
		[this.value1,this.value2] = [+prompt('value1?',0),+prompt('value2?',0)];
	}

	this.sum = function() {
		return this.value1 + this.value2;
	}

	this.mul = function() {
		return this.value1 * this.value2;
	}
}

let mao = new Calculator();
mao.read();