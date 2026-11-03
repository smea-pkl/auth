import { useLocation, useParams, useNavigate } from "react-router-dom";

export default function Room() {
  const { roomID } = useParams();
  const location = useLocation();
  const navigate = useNavigate();

  const opponent = location.state?.opponent || "Lawan tidak diketahui";
  const currentUser = location.state?.currentUser || "Kamu";

  return (
    <div>
      <h2>Room Pertandingan</h2>
      <p>
        ID Room: <strong>{roomID}</strong>
      </p>

      <div>
        <h3>
          {currentUser} VS {opponent}
        </h3>
      </div>

      <button onClick={() => navigate("/home")}>Tinggalkan Room</button>
    </div>
  );
}
