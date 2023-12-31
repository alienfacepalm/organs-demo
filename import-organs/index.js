const admin = require('firebase-admin');
const serviceAccount = require('../organs-demo-api-key.json');
const csv = require('csv-parser');
const fs = require('fs');

admin.initializeApp({
  credential: admin.credential.cert(serviceAccount)
});

const db = admin.firestore();

fs.createReadStream('organs.csv')
  .pipe(csv())
  .on('data', (row) => {
    db.collection('organs').add(row);
  })
  .on('end', () => {
    console.log('CSV file successfully processed');
  });
