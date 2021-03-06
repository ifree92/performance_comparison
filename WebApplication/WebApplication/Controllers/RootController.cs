using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using WebApplication.DTO;
using WebApplication.Services;

namespace WebApplication.Controllers
{
    [ApiController]
    [Route("/")]
    public class RootController : ControllerBase
    {
        private UsersService _usersService;

        public RootController(UsersService usersService)
        {
            _usersService = usersService;
        }

        [HttpGet("")]
        public string GetRoot()
        {
            return "Root";
        }

        [HttpGet("string")]
        public string GetString()
        {
            return "Hello world!";
        }

        [HttpGet("json")]
        public UserDTO GetJson()
        {
            UserDTO userDto = _usersService.GetUser();
            return _usersService.GetUser();
        }

        [HttpGet("json-array")]
        public List<UserDTO> GetJsonArray()
        {
            return _usersService.GetUsers();
        }

        [HttpGet("json-array-task")]
        public Task<List<UserDTO>> GetJsonArrayTask()
        {
            return _usersService.GetUsersTask();
        }
    }
}