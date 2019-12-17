import express from 'express';
import { appRouter } from './app.router';

const app = express();

async function main() {
  app.use('/', appRouter);
  app.listen(7777, () => console.log('Listening 7777'));
}

main().catch(console.error);
