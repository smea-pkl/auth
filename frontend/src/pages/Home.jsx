import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Home() {
  const navigate = useNavigate();
  const [pesan, setPesan] = useState("Memuat data...");
  const [isSearching, setIsSearching] = useState(false);
  const [username, setUsername] = useState("");

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      navigate("/login");
      return;
    }

    fetch("http://localhost:8080/profile", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => {
        if (!res.ok) throw new Error("Token tidak valid");
        return res.text();
      })
      .then((data) => {
        setPesan(`Siap bermain, ${data}`);
        setUsername(data);
      })
      .catch(() => {
        localStorage.removeItem("token");
        navigate("/login");
      });
  }, [navigate]);

  const startMatchmaking = () => {
    setIsSearching(true);
    const ws = new WebSocket(
      `ws://localhost:8080/ws/matchmake?user=${username}`,
    );

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.type == "MATCH_FOUND") {
        ws.close();
        navigate(`/room/${data.roomID}`, {
          state: { opponent: data.opponent, currentUser: username },
        });
      }
    };

    ws.onclose = () => setIsSearching(false);
  };

  return (
    <div>
      <h2>Beranda</h2>
      <p>Status: {pesan}</p>

      <div>
        <button onClick={startMatchmaking} disabled={isSearching}>
          {isSearching ? "Mencari Lawan..." : "Cari Lawan"}
        </button>

        <button
          onClick={() => {
            localStorage.removeItem("token");
            navigate("/login");
          }}
        >
          Logout
        </button>
      </div>
    </div>
  );
}
