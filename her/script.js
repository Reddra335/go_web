// Fungsi untuk Update Jam dan Tanggal
function updateTime() {
    const timeElement = document.getElementById('time');
    const dateElement = document.getElementById('date');
    const greetingElement = document.getElementById('greeting');

    const now = new Date();

    // Format Jam (HH:MM)
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    timeElement.textContent = `${hours}:${minutes}`;

    // Format Tanggal
    const options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };
    dateElement.textContent = now.toLocaleDateString('id-ID', options);

    // Logika Ucapan berdasarkan waktu
    const currentHour = now.getHours();
    let greeting = 'Gorengan Rek,';

    if (currentHour >= 5 && currentHour < 12) {
        greeting = 'Selamat Pagi,';
    } else if (currentHour >= 12 && currentHour < 15) {
        greeting = 'Selamat Siang,';
    } else if (currentHour >= 15 && currentHour < 18) {
        greeting = 'Selamat Sore,';
    }

    greetingElement.textContent = greeting;
}

// Data Quotes Random (Bisa kamu isi dengan kata-kata keren/motivasi)
const quotes = [
    { text: "Jangan lupa istirahat, sistem juga butuh cooldown.", author: "UI System" },
    { text: "Kode yang bagus adalah kode yang jalan tanpa error. Hidup juga gitu.", author: "Admin" },
    { text: "Tetap fokus. Masa depan diciptakan dari apa yang kamu lakukan hari ini.", author: "System Log" },
    { text: "Error 404: Motivasi not found. Silakan minum kopi dulu.", author: "Coffee.js" },
    { text: "Semoga harimu seindah UI design tanpa bug.", author: "Developer" }
];

function generateQuote() {
    const quoteText = document.getElementById('quote');
    const quoteAuthor = document.querySelector('.author');

    const randomIndex = Math.floor(Math.random() * quotes.length);
    const randomQuote = quotes[randomIndex];

    quoteText.textContent = `"${randomQuote.text}"`;
    quoteAuthor.textContent = `- ${randomQuote.author}`;
}

// Menjalankan fungsi saat web dibuka
updateTime();
generateQuote();

// Update jam setiap 1 detik
setInterval(updateTime, 1000);