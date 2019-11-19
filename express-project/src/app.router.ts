import { Router, Request, Response } from 'express';
import { appService } from './app.service';

export const appRouter = Router();

appRouter.get('/string', (req: Request, res: Response) => {
  res.json('Hello world!');
});

appRouter.get('/json', (req: Request, res: Response) => {
  res.json(appService.getUser());
});

appRouter.get('/json-array', (req: Request, res: Response) => {
  res.json(appService.getUsers());
});

appRouter.get('/validated-users', (req: Request, res: Response) => {});
