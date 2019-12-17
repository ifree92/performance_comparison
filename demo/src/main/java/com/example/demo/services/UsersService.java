package com.example.demo.services;

import com.example.demo.dto.UserDTO;
import org.springframework.stereotype.Service;

import java.util.ArrayList;

@Service
public class UsersService {
    private UserDTO userDTO;
    private ArrayList<UserDTO> userDTOList = new ArrayList<>();

    public UsersService() {
        userDTO = new UserDTO();
        userDTO.id = 1;
        userDTO.firstName = "Hello";
        userDTO.lastName = "World";
        userDTO.age = 44;
        userDTO.hobbies = new String[]{"football", "basketball", "guitar", "biking"};

        for (int i = 0; i < 1000; i++) {
            UserDTO usr = new UserDTO();
            usr.id = 1;
            usr.firstName = "Hello";
            usr.lastName = "World";
            usr.age = 44;
            usr.hobbies = new String[]{"football", "basketball", "guitar", "biking"};
            userDTOList.add(usr);
        }
    }

    public UserDTO getUser() {
        return userDTO;
    }

    public ArrayList<UserDTO> getUsers() {
        return this.userDTOList;
    }
}
