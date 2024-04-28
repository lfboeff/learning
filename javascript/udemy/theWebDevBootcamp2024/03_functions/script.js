// # our very first function:
// function singSong() {
//     console.log("DO");
//     console.log("RE");
//     console.log("MI");
// }
// singSong();


// # arguments intro:
// function greet(firstName) {
//     console.log(`Hello, ${firstName}!`);
// }
// greet("Natália");


// # multiple arguments:
// function greet(firstName, lastName) {
//     console.log(`Hey there, ${firstName} ${lastName}!`)
// }
// greet("Tio", "João");

// function repeat(word, num) {
//     finalWord = "";
//     for(let i = 0; i < parseInt(num); i++) {
//         finalWord += word;
//     }
//     console.log(finalWord);
// }
// repeat("hi", 3);


// # the return keyword:
function add(num1, num2) {
    return num1+num2
}
let sum = add(5,5);
console.log(sum);