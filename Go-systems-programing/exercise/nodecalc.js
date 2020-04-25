#!/usr/bin/env node

if(process.argv.length <=  2){
    console.log(`Usage: ./file.js "[number] <operation> [number]"`);
    console.log(`	Example : ./nodecalc.js "(12 * 45  / 565) " \n`);
    process.exit(1)
}

operation =  process.argv.slice(2)

let op = ""

for(let i = 0; i < operation.length; i++){
    op = ' ' + operation[i]
}

console.log(`Answer is : ${eval(op)}`)

//console.log(op)
