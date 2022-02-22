const fs = require('fs');
const { exec } = require('child_process');

const begin = new Date();
const fileName = `${begin.getDay()}${begin.getMonth()}${begin.getFullYear()}${begin.getHours()}${begin.getMinutes()}${begin.getSeconds()}`;
console.log(`---------------------------------\nExecution started @${begin.getHours()}:${begin.getMinutes()}:${begin.getSeconds()}`);
exec(`papertrail \'received response for https://api.tap.company/v2/charges\' --min-time \'2022-02-19 08:00:00\' --max-time \'2022-02-21 08:00:00\' > ${fileName}.txt`, (error,output,stderr) => {
    console.log('Command executed.........');
    const end = new Date();
    console.log(`Execution completed @${end.getHours()}:${end.getMinutes()}:${end.getSeconds()}\n---------------------------------`);
    if(error||stderr){
        console.log('Something went wrong');
    }else{
        console.log('Reading file.......\n---------------------------------')
        fs.readFile(`${fileName}.txt`, 'utf8',(error, data) => {
            if(error){
                console.log('Failed to read file');
            }else{
                const extracted = data.match(/([0-9]+\.[0-9]+s)/g);
        
                let timeInFloat = extracted.map((timeInString) => {return parseFloat(timeInString)});
                console.log(`Total : ${timeInFloat.length} records found`);
                console.log(`Minimun : ${Math.min(...timeInFloat)}`);
                const maxValue = Math.max(...timeInFloat);
                console.log(`Maximun : ${maxValue}`);

                var sum = 0;
                timeInFloat.forEach((time) => {sum +=time;});

                console.log(`Average : ${sum/timeInFloat.length}`);

                //Sorting times to get value(s) at/from specific pecentile position
                timeInFloat.sort();
                console.log(`95th Percentile : ${timeInFloat[Math.round((95*timeInFloat.length)/100)]}`);
                console.log(`99th Percentile : ${timeInFloat[Math.round((99*timeInFloat.length)/100)]}`);
                console.log('---------------------------------');
            }
        });        
    }
});