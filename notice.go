package jpushv3

type Notification struct {
	Alert   string                 `json:"alert,omitempty"`
	Android map[string]interface{} `json:"android,omitempty"`
	Ios     map[string]interface{} `json:"ios,omitempty"`
}

func NewNotification(alert string) *Notification {
	n := Notification{}
	n.Alert = alert
	return &n
}

func (this *Notification) SetSimpleAlert(alert string) {
	this.Alert = alert
}

func (this *Notification) initAndroid() {
	if this.Android == nil {
		this.Android = map[string]interface{}{}
	}
}

func (this *Notification) initIOS() {
	if this.Ios == nil {
		this.Ios = map[string]interface{}{}
	}
}

func (this *Notification) SetAndroidAlert(alert string) {
	this.initAndroid()
	this.Android["alert"] = alert
}

func (this *Notification) SetAndroidTitle(title string) {
	this.initAndroid()
	this.Android["title"] = title
}

func (this *Notification) SetIOSAlert(alert string) {
	this.initIOS()
	this.Ios["alert"] = alert
}

func (this *Notification) SetIOSSound(sound string) {
	this.initIOS()
	this.Ios["sound"] = sound
}

func (this *Notification) SetIOSBadge(badge string) {
	this.initIOS()
	this.Ios["badge"] = badge
}

func (this *Notification) SetIOSExtras(extras interface{}) {
	this.initIOS()
	this.Ios["extras"] = extras
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
