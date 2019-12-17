package com.example.demo;

import com.example.demo.dto.UserDTO;
import com.example.demo.services.UsersService;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class RootController {
    private UsersService usersService;

    public RootController(UsersService usersService) {
        this.usersService = usersService;
    }

    @RequestMapping(value = "/", method = RequestMethod.GET)
    public String index() {
        return "Root";
    }

    @RequestMapping(value = "/string", method = RequestMethod.GET)
    public String string() {
        return "Hello world!";
    }

    @RequestMapping(value = "/json", method = RequestMethod.GET)
    public UserDTO getJson() {
        return this.usersService.getUser();
    }

    @RequestMapping(value = "/json-array", method = RequestMethod.GET)
    public List<UserDTO> getJsonArray() {
        return this.usersService.getUsers();
    }
}
