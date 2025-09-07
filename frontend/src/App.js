import React, { useEffect, useState } from "react";
import "./App.css"; // import file css

export default function App() {
  const [stations, setStations] = useState([]);
  const [selectedStation, setSelectedStation] = useState(null);
  const [schedule, setSchedule] = useState([]);
  const [stationFilter, setStationFilter] = useState("");
  const [scheduleFilter, setScheduleFilter] = useState("");

  // Ambil daftar stasiun
  useEffect(() => {
    fetch("http://localhost:8080/v1/api/stations/")
      .then((res) => res.json())
      .then((data) => {
        if (data.success) {
          setStations(data.data);
        }
      })
      .catch((err) => console.error("Error fetch stations:", err));
  }, []);

  // Ambil jadwal saat stasiun dipilih
  const fetchSchedule = (id) => {
    fetch(`http://localhost:8080/v1/api/stations/${id}`)
      .then((res) => res.json())
      .then((data) => {
        if (data.success) {
          setSchedule(data.data);
          setSelectedStation(id);
          setScheduleFilter(""); // reset filter jadwal saat ganti stasiun
        } else {
          setSchedule([]);
        }
      })
      .catch((err) => console.error("Error fetch schedule:", err));
  };

  return (
    <div className="container">
      {/* Title Box */}
      <h1 className="title-box">
        🚇 Daftar Stasiun MRT Jakarta <span className="count">({stations.length})</span>
      </h1>

      {/* Filter stasiun */}
      <div className="filter-box">
        <input
          type="text"
          placeholder="Cari stasiun..."
          value={stationFilter}
          onChange={(e) => setStationFilter(e.target.value)}
        />
      </div>

      {/* Grid stasiun */}
      <div className="station-grid">
        {stations
          .filter((station) =>
            station.name.toLowerCase().includes(stationFilter.toLowerCase())
          )
          .map((station) => (
            <button
              key={station.id}
              onClick={() => fetchSchedule(station.id)}
              className={`station-btn ${
                selectedStation === station.id ? "active" : ""
              }`}
            >
              <p className="station-title">{station.name}</p>
              <span className="station-id">ID: {station.id}</span>
            </button>
          ))}
      </div>

      {/* Jadwal */}
      <div className="schedule-box">
        <h2>📅 Jadwal Keberangkatan</h2>

        {/* Filter jadwal */}
        {selectedStation && schedule.length > 0 && (
          <div className="filter-box">
            <input
              type="text"
              placeholder="Cari jadwal..."
              value={scheduleFilter}
              onChange={(e) => setScheduleFilter(e.target.value)}
            />
          </div>
        )}

        {selectedStation && schedule.length === 0 && (
          <p className="empty">Tidak ada jadwal untuk stasiun ini</p>
        )}

        {!selectedStation && (
          <p className="empty">Pilih stasiun untuk melihat jadwal</p>
        )}

        <ul>
          {schedule
            .filter(
              (item) =>
                item.stationName
                  .toLowerCase()
                  .includes(scheduleFilter.toLowerCase()) ||
                item.time.includes(scheduleFilter)
            )
            .map((item, index) => (
              <li key={index} className="schedule-item">
                <span className="station-name">{item.stationName}</span> —{" "}
                <span className="time">{item.time}</span>
              </li>
            ))}
        </ul>
      </div>
    </div>
  );
}
