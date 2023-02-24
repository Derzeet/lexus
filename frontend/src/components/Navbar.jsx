import { Link } from 'react-router-dom';


function LogoutButton() {
    function handleLogout() {
      localStorage.removeItem('token'); // remove token from localStorage
      window.location.reload();
    }
  
    return (
      <button onClick={handleLogout}>
        Logout
      </button>
    );
  }
function Navbar() {
    const  token  = localStorage.getItem('token');
    const CHECK = !!token
    
  return (

    <nav>
      <ul>
        {!CHECK && (
            <>
            <li>
            <Link to="/">Login</Link>
          </li>
          <li>
            <Link to="/register">Register</Link>
          </li>
            </>
        )}

        <li>
          <Link to="/store">Store</Link>
        </li>
        {CHECK && (
          <>
            <li>
              <Link to="/profile">Profile</Link>
            </li>
            <li>
              <Link to="/creategun">Create</Link>
            </li>
            <li>
              <LogoutButton />
            </li>
          </>
        )}
      </ul>
    </nav>
  );
}

export default Navbar;
