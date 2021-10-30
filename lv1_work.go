package main


type Author struct {
	name string;
	Focus int64;
	VIP bool;
	video_number int64;
	dianzan int;
	toubi int;


}

func main() {

	var bidao Author;
	bidao.name="毕导THU";
	bidao.Focus=45950000
	bidao.VIP=true;
	bidao.video_number=124
	print("name:\n",bidao.name)
	print("粉丝数：",bidao.Focus)


}

