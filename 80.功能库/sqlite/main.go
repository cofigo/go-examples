package main
import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)
func main() {
	homeDir, _ := os.UserHomeDir()
	db, err := sql.Open("sqlite3", homeDir+"/Library/Containers/com.netease.163music/Data/Documents/storage/sqlite_storage.sqlite3")
	checkErr(err)
	//插入数据
	//stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	//checkErr(err)
	//res, err := stmt.Exec("361way", "研发部", "2019-03-06")
	//checkErr(err)
	//id, err := res.LastInsertId()
	//checkErr(err)
	//fmt.Println(id)
	//更新数据
	//stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	//checkErr(err)
	//res, err = stmt.Exec("361wayupdate", id)
	//checkErr(err)
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//fmt.Println(affect)
	//查询数据
	rows, err := db.Query("SELECT id, type, type_extra, track_id, detail, size, bitrate, state, download_time, complete_time, album_id, relative_path, track_name, artist_name, album_name, id3_state, real_suffix, file_type FROM web_offline_track")
	//rows, err := db.Query("SELECT id, type, type_extra, track_id, detail, size, bitrate, state, download_time, complete_time, album_id, relative_path, track_name, artist_name, album_name, id3_state, real_suffix, file_type FROM web_offline_track  where relative_path!=\"\"")
	//rows, err := db.Query("SELECT id, type, track_id, detail, size, bitrate, state, download_time, complete_time, album_id, relative_path, track_name, artist_name, album_name, id3_state, real_suffix, file_type FROM web_offline_track where relative_path=?", "/天下 - 张杰.ncm")

	checkErr(err)
	for rows.Next() {
		var id string
		var type_id int
		type_extra := sql.NullString{
			String: "",
			Valid:  false,
		}
		var track_id int
		var detail string
		var size int
		//dfsid := sql.NullInt64{Int64: 0, Valid: false,}
		var bitrate int
		var state int
		var download_time int
		var complete_time int
		//source_href := sql.NullString{String: "", Valid:  false,}
		//source_text := sql.NullString{String: "", Valid:  false,}
		//source_extra := sql.NullString{String: "", Valid:  false,}
		var album_id string
		var relative_path string
		var track_name string
		var artist_name string
		var album_name string
		var id3_state int
		var real_suffix string
		var file_type int

		err = rows.Scan(&id, &type_id, &type_extra, &track_id, &detail, &size, &bitrate, &state, &download_time, &complete_time, &album_id, &relative_path, &track_name, &artist_name, &album_name, &id3_state, &real_suffix, &file_type)
		checkErr(err)
		//fmt.Println(id)
		//fmt.Println(type_id)
		//fmt.Println(type_extra)
		//fmt.Println(track_id)
		//fmt.Println(detail)
		//fmt.Println(size)
		//fmt.Println(dfsid)
		//fmt.Println(bitrate)
		//fmt.Println(state)
		//fmt.Println(download_time)
		//fmt.Println(complete_time)
		//fmt.Println(source_href)
		//fmt.Println(source_text)
		//fmt.Println(source_extra)
		//fmt.Println(album_id)
		//fmt.Println(relative_path)
		//fmt.Println(track_name)
		//fmt.Println(artist_name)
		//fmt.Println(album_name)
		//fmt.Println(id3_state)
		//fmt.Println(real_suffix)
		//fmt.Println(file_type)
		fmt.Println("ID: ", track_id, " IMG: ", id)
		fmt.Println("Track: ", track_name, " Artist: ", artist_name, " Album: ", album_name)
		fmt.Println("Type: ", file_type ," File: ", relative_path)
		fmt.Println("")
	}
	//删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//res, err = stmt.Exec(id)
	//checkErr(err)
	//affect, err = res.RowsAffected()
	//checkErr(err)
	//fmt.Println(affect)
	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
