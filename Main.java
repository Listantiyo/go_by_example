import java.util.Scanner;

class Kamar {
    int nomor;
    String tipe;
    double harga;
    boolean tersedia;

    public Kamar(int nomor, String tipe, double harga, boolean tersedia) {
        this.nomor = nomor;
        this.tipe = tipe;
        this.harga = harga;
        this.tersedia = tersedia;
    }
}

public class Main {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        // daftar kamar
        Kamar room1 = new Kamar(101, "Standar", 200000, true);
        Kamar room2 = new Kamar(102, "Superior", 250000, true);
        Kamar room3 = new Kamar(103, "Duluxe", 300000, true);
        Kamar room4 = new Kamar(104, "Suite", 350000, true);

        Kamar[] daftarKamar = {room1, room2, room3, room4};

        // tampilan menu
        System.out.println("Daftar kamar Tersedia:");
        System.out.println("......................");
        System.out.println("No  | Tipe     | Harga/malam | Tersedia");
        System.out.println("........................................");

        for (Kamar k : daftarKamar) {
            System.out.println(k.nomor + "  | " + k.tipe + " | " + k.harga + "     | " + k.tersedia);
        }

        System.out.println("\nSilahkan masukan data pemesanan (maksimal 3 kamar!)");

        Kamar kamar1 = null, kamar2 = null, kamar3 = null;
        int malam1 = 0, malam2 = 0, malam3 = 0;

        System.out.print("Masukkan nomor kamar 1 (atau 0 untuk skip): ");
        int no1 = input.nextInt();
        if (no1 != 0) {
            kamar1 = cariKamar(daftarKamar, no1);
            if (kamar1 != null) {
                System.out.print("Lama menginap kamar 1: ");
                malam1 = input.nextInt();
            }
        }

        System.out.print("Masukkan nomor kamar 2 (atau 0 untuk skip): ");
        int no2 = input.nextInt();
        if (no2 != 0) {
            kamar2 = cariKamar(daftarKamar, no2);
            if (kamar2 != null) {
                System.out.print("Lama menginap kamar 2: ");
                malam2 = input.nextInt();
            }
        }

        System.out.print("Masukkan nomor kamar 3 (atau 0 untuk skip): ");
        int no3 = input.nextInt();
        if (no3 != 0) {
            kamar3 = cariKamar(daftarKamar, no3);
            if (kamar3 != null) {
                System.out.print("Lama menginap kamar 3: ");
                malam3 = input.nextInt();
            }
        }

        // hitung total
        double total = 0, layanan = 0, pajak = 0;

        if (kamar1 != null) {
            total += kamar1.harga * malam1;
            layanan += 50000;
        }
        if (kamar2 != null) {
            total += kamar2.harga * malam2;
            layanan += 50000;
        }
        if (kamar3 != null) {
            total += kamar3.harga * malam3;
            layanan += 50000;
        }

        pajak = total * 0.1;
        double totalSebelumDiskon = total + layanan + pajak;
        double diskon = 0;

        if (total > 500000) {
            diskon = totalSebelumDiskon * 0.15;
        }

        double totalAkhir = totalSebelumDiskon - diskon;

        // cetak struk
        System.out.println("\n......... Struk Reservasi .........");
        if (kamar1 != null) cetakKamar(kamar1, malam1);
        if (kamar2 != null) cetakKamar(kamar2, malam2);
        if (kamar3 != null) cetakKamar(kamar3, malam3);

        System.out.println(".....................................");
        System.out.println("Subtotal: Rp " + total);
        System.out.println("Pajak (10%): Rp " + pajak);
        System.out.println("Biaya layanan: Rp " + layanan);

        if (diskon > 0) {
            System.out.println("Diskon 15%: Rp " + diskon);
        }
        System.out.println("Total Bayar: Rp " + totalAkhir);

        if (total > 300000) {
            System.out.println("Anda menerima bonus gratis sarapan!");
        }
        System.out.println(".....................................");

        input.close();
    }

    static Kamar cariKamar(Kamar[] list, int nomor) {
        for (Kamar k : list) {
            if (k.nomor == nomor) return k;
        }
        return null;
    }

    static void cetakKamar(Kamar kamar, int malam) {
        double totalTagihan = kamar.harga * malam;
        System.out.println("Kamar " + kamar.nomor + " dengan tipe: " + kamar.tipe + " menginap selama: " + malam + " malam");
        System.out.println("Harga/malam: Rp " + kamar.harga + " | Total: Rp " + totalTagihan);
    }
}
