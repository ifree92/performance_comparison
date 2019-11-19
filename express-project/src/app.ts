import express from 'express';
import { appRouter } from './app.router';

const app = express();

async function main() {
  app.use('/', appRouter);
  app.listen(5555, () => console.log('Listening 5555'));
}

main().catch(console.error);
