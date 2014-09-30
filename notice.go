package jpushv3


type Notification struct {
	values map[string]interface {}
	android map[string]interface {}
	ios map[string]interface {}
}

func (this *Notification) SetSimpleAlert(alert string) {
	this.values["alert"] = alert
}


func (this *Notification) initAndroid() {
	if this.android == nil {
		this.android = map[string]interface {}{}
		this.values["android"] = this.android
	}
}

func (this *Notification) initIOS() {
	if this.ios == nil {
		this.ios = map[string]interface {}{}
		this.values["ios"] = this.ios
	}
}



func (this *Notification) SetAndroidAlert(alert string) {
	this.initAndroid()
	this.android["alert"] = alert
}


func (this *Notification) SetAndroidTitle(title string) {
	this.initAndroid()
	this.android["title"] = title
}

func (this *Notification) SetIOSAlert(alert string) {
	this.initIOS()
	this.ios["alert"] = alert
}

func (this *Notification) SetIOSSound(sound string) {
	this.initIOS()
	this.ios["sound"] = sound
}

func (this *Notification) SetIOSBadge(badge string) {
	this.initIOS()
	this.ios["badge"] = badge
}


type Notice struct {
	Alert string `json:"alert"`
}

type AndroidNotice struct {
	Object NoticeAndroid `json:"android"`
}

type IosNotice struct {

}

type NoticeAndroid struct {
	Alert     string `json:"alert"`
	Title     string `json:"title"`
	BuilderId int    `json:"builder_id"`
	//Extras    map[string]string `json:"extras"`
}

func (this *AndroidNotice) SetAlert(alert string) {
	this.Object.Alert = alert

}

func (this *AndroidNotice) SetTitle(title string) {
	this.Object.Title = title
}

func (this *AndroidNotice) SetBuilderId(id int) {
	this.Object.BuilderId = id
}

//func (this *AndroidNotice) AddExtras(key, value string) {
//	if this.notice.Extras == nil {
//		this.notice.Extras = make(map[string]string)
//	}
//	this.notice.Extras[key] = value
//}
