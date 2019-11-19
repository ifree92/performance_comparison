export interface IUser {
  id: number;
  firstName: string;
  lastName: string;
  age: number;
  hobbies: string[];
}

class AppService {
  private user: IUser = {
    id: 1,
    firstName: "Hello",
    lastName: "World",
    age: 44,
    hobbies: ["football", "basketball", "guitar", "biking"]
  };

  private users: IUser[] = [];

  constructor() {
    for (let i = 0; i < 1000; i++) {
      this.users.push({ ...this.user });
    }
  }

  getUser() {
    return this.user;
  }

  getUsers() {
    return this.users;
  }
}

export const appService = new AppService();
