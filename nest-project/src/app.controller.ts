import { Controller, Get } from '@nestjs/common';
import { AppService, IUser } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('string')
  getString(): string {
    return 'Hello world!';
  }

  @Get('json')
  getJson(): IUser {
    return this.appService.getUser();
  }

  @Get('json-array')
  getJsonArray(): IUser[] {
    return this.appService.getUsers();
  }

  @Get('validated-users')
  getValidatedUsers() {
    return '';
  }
}
