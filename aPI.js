const https = require('https');
const fs = require('fs');
require('dotenv').config()

console.log('Checking compilation')

const options = {

    headers: {
        'X-Papertrail-Token' : process.env.PPTRL_API_KEY,
    }
}

const req = https.get('https://papertrailapp.com/api/v1/events/search.json?q=error', options,res => {
    console.log(`statusCode: ${res.statusCode}`)
  
    res.on('data', d => {
    //   process.stdout.write(d)
      fs.writeFile('testWrite.txt',d, error => {
          if(error){
              console.log(error);
          }
      })

    })
  })
  
  req.on('error', error => {
    console.error(`There has been a error : ${error}`)
  })