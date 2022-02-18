const fs = require('fs');
const { exec } = require('child_process');

exec('papertrail \'error tap gateway timed out\' --min-time \'2022-02-10 00:00:00\' --max-time \'2022-02-16 00:00:00\' > testPapertrail.txt ', (error,output,stderr) => {
    if(error||stderr){
        console.log('Something went wrong');
    }else{
        fs.readFile('testPapertrail.txt', 'utf8',(error, data) => {
            if(error){
                console.log('Failed to read file');
            }else{
                let extracted = data.match(/([0-9]+\.[0-9]+s)/g);
        
                let timeInFloat = extracted.map((timeInString) => {return parseFloat(timeInString)});
                console.log(`Total : ${timeInFloat.length} records found`);
                console.log(`Minimun : ${Math.min(...timeInFloat)}`);
                console.log(`Maximun : ${Math.max(...timeInFloat)}`)
            }
        });        
    }
});


// fs.readFile('/Users/harsh/testPapertrail.txt', 'utf8',(error, data) => {
//     if(error){
//         console.log('Failed to read file');
//     }else{
//         let extracted = data.match(/([0-9]+\.[0-9]+s)/g);

//         let timeInFloat = extracted.map((timeInString) => {return parseFloat(timeInString)});
//         console.log(`Total : ${timeInFloat.length} records found`);
//         console.log(`Minimun : ${Math.min(...timeInFloat)}`);
//         console.log(`Maximun : ${Math.max(...timeInFloat)}`)
//     }
// });