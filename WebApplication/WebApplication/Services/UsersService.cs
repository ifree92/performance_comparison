using System.Collections.Generic;
using System.Threading.Tasks;
using WebApplication.DTO;

namespace WebApplication.Services
{
    public class UsersService
    {
        private UserDTO _user;
        private List<UserDTO> _users = new List<UserDTO>();

        public UsersService()
        {
            _user = new UserDTO();
            _user.Id = 1;
            _user.FirstName = "Hello";
            _user.LastName = "World";
            _user.Age = 44;
            _user.Hobbies = new[] {"football", "basketball", "guitar", "biking"};


            for (int i = 0; i < 1000; i++)
            {
                UserDTO userDto = new UserDTO();
                userDto.Id = 1;
                userDto.FirstName = "Hello";
                userDto.LastName = "World";
                userDto.Age = 44;
                userDto.Hobbies = new[] {"football", "basketball", "guitar", "biking"};
                _users.Add(userDto);
            }
        }

        public UserDTO GetUser()
        {
            return _user;
        }

        public List<UserDTO> GetUsers()
        {
            return _users;
        }

        public Task<List<UserDTO>> GetUsersTask()
        {
            return Task.FromResult(_users);
        }
    }
}