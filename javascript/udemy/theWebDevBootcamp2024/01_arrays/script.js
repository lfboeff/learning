// push & pop / shift & unshift:
// let days = ['Monday', 'Tuesday', 'Wednesday'];
// console.log(days);
// days.push("Thursday");
// console.log(days);
// days.pop();
// console.log(days);
// days.shift();
// console.log(days);
// days.unshift("Monday");
// console.log(days);

// concat, indexOf, includes & reverse:
// cats = ["blue", "kitty"];
// dogs = ["rusty", "wyatt"];
// let comboParty = cats.concat(dogs);
// console.log(comboParty);
// console.log(cats.includes("blue"));
// console.log(cats.includes("Blue"));
// console.log(comboParty.indexOf("rusty"));
// console.log(comboParty.indexOf("toby"));
// comboParty.reverse();
// console.log(comboParty);

// slice & splice:
// let colors = ["red", "orange", "yellow", "green", "blue", "indigo", "violet"];
// console.log(colors);
// console.log(colors.slice(0,3));
// console.log(colors.slice(-1));
// colors.splice(5, 1, "DELETED");
// console.log(colors);

// references types & equality testing:
// let number1 = 90;
// let number2 = 90;
// console.log(number1 === number2);
// let arr = [1, 2, 3];
// console.log(arr === arr);

// arrays + const:
// const nums = [1, 2, 3]; // as long as we keep the array reference the same, we can modify the array
// nums.push(4);
// console.log(nums);

// multi-dimensional arrays:
const colors = [
    ["red", "crimson"],
    [ "green", "olive"],
    ["blue", "navy blue"],
];
console.log(colors);
const board =[
    ["0", null, "X"],
    [null, "X", "0"],
    ["X", "0", null],
];
console.log(board);
const gameBoard = [
    ["X", "0", "X"],
    ["0", null, "X"],
    ["0", "0", "X"],
];
console.log(gameBoard);
console.log(gameBoard[0][0]);
console.log(gameBoard[1][1]);
console.log(gameBoard[1][2]);
console.log(gameBoard[2][0]);
