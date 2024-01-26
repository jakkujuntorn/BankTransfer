- จำลองเว็บไซด์  ธนาคาร (ฝาก, โอน, ถอน, เรียกดู statement)
 
        -gRPC (Gateway)
        -Hexagonal Architecture
            -port and adaptors
        -DB
            -Postgres 
        -Code Structure
            -By Feature
        
- middleware gRPC ขอข้ามไปก่อน มันเยอะสะเหลือเกิน

- Valisate data ต่างๆ
    -ทำเฉพาะ Transfer  เช็ค money ติดลบ กับ 0
- จะใส ECHO เพิ่ม เพื่อไม่ใช้ gRPC Gateway ***********

-API ทุกเส้น ต้องปั้น response สวยๆออกไป
- API User
  
        -Login เทสแล้ว
        -CreateUser ยังไม่ได้เทส **************
        -GetUser_ByUsername เทสแล้ว 
        -UpdateUser เทสแล้ว
        -ChangePassword  เทสแล้ว

- API Account

        -CreateAccount - ยังไม่ได้เทส เพราะ อยากให้ query เช็คสกุลเงินเลย
                     ไม่อยากให้ดึงข้อมูลมาเช็คใน golang ***
                    - return
        -GetAccount เทสแล้ว
            -ปรับ ปั้นข้อมูลตรง time
        -ListAccount เทสแล้ว
        -Delete ยังไม่ได้เทส

- API Transfers

        -Create Transfers เทสแล้ว
            -ไม่สาามรถใช้ struct ของ proto เข้าไป create ใน  DB ได้ ต้องเป็น struct ปกติของ go
                -หรือเราไม่สาารถทำได้หว่า ***
        -Create_Deposit เทสแล้ว
        -Create_Withdraw เทสแล้ว
        -GetTransfer_ById เทสแล้ว
        -GetTransfer_ByOwner เทสแล้ว
