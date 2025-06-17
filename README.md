# 🛠️ File Organizer CLI - Usage Guide

A command-line tool to organize files in a directory based on file extension.

---

## 🎮 Installation (Local)

You can run it directly using Go:

```bash
go run main.go [command] --dir path/to/folder
```

Or build the binary:

```bash
go build -o organizer main.go
./organizer [command] --dir path/to/folder
```

If you'd like to install it globally:

```bash
go build -o organizer main.go
sudo mv organizer /usr/local/bin/
```

Now you can run it anywhere:

```bash
organizer run --dir ~/Downloads
```

---

## ⚙️ Commands

### 1. `run`

📆 จัดเรียงไฟล์จริงตามกฎใน `rules.json`

```bash
go run main.go run --dir ~/Downloads --rules rules.json
```

📌 **Options:**

* `--dir`, `-d` → โฟลเดอร์เป้าหมาย (จำเป็น)
* `--rules`, `-r` → path ไปยัง rules.json (ค่า default: `rules.json`)
* `--interactive`, `-i` → โหมดยืนยันก่อย้าย แต่ละ file

#### ตัวอย่าง:

```bash
go run main.go run --dir ~/Downloads --interactive
```

หรือ ถ้ามีชื่อ folder 2พยางค์

```bash
go run main.go run --dir ~/Downloads/Telegram\ Desktop --interactive
```

---

### 2. `preview`

📅 แสดงตัวอย่างไฟล์ที่จะถึงย้าย โดยไม่ย้ายจริง

```bash
go run main.go preview --dir ~/Downloads --rules rules.json
```

หรือ ถ้ามีชื่อ folder 2พยางค์

```bash
go run main.go preview --dir ~/Downloads/Telegram\ Desktop
```

🔹 เหมาะสำหรับตรวจสอบก่อน run จริง

#### ตัวอย่าง:

```bash
go run main.go preview --dir ~/Downloads
```

---

### 3. `revert`

🔄 ยกเลิกการจัดระเบียบไฟล์ครั้งล่าสุด โดยอ้างอิงจาก `history/moves.json`

```bash
go run main.go revert
```

⚠️ ไฟล์จะถึงย้ายกลับตำแหน่งเดิม ถ้า history ยังไม่ถึง

#### ตัวอย่าง:

```bash
go run main.go revert
```

---

## 📜 ตัวอย่าง `rules.json`

```json
{
  ".jpg": "Images",
  ".png": "Images",
  ".txt": "Text",
  ".pdf": "PDFs",
  ".mp3": "Music",
  ".zip": "Zip",
  ".xlsx": "Xls",
  ".docx": "Doc"
}
```

---

## 📃 ตัวอย่างการใช้งาน

### จัดเรียงไฟล์ในโฟลเดอร์ `Downloads` ทันที

```bash
go run main.go run --dir ~/Downloads
```

หรือ

### ใช้ interactive mode เพื่อยืนยัน Y/N

```bash
go run main.go run --dir ~/Downloads --interactive
```

### แสดง preview ก่อน run

```bash
go run main.go preview --dir ~/Downloads
```

หรือ ถ้ามีชื่อ folder 2พยางค์

```bash
go run main.go preview --dir ~/Downloads/Telegram\ Desktop
```

### ย้อนกลับการย้ายไฟล์

```bash
go run main.go revert
```

---

## 📌 Tips

* ควร run `preview` ก่อ `run` เสมอเพื่อความปลอดภัย
* ปรับแต่ง `rules.json` ได้ตามต้องการ เช่น `.docx`, `.xlsx`, `.zip`

---

## 🥪 Testing
****
แนะนำให้ลองใช้กับโฟลเดอร์ทดสอบก่อน เพื่อป้องกันการสูญเสียไฟล์สำคัญ

---
