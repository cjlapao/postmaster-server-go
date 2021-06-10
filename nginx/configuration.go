package nginx

type Configuration struct {
	User               string
	WorkerProcesses    string
	ErrorLog           string
	AccessLog          string
	WorkerRlimitNoFile string
	PID                string
	// Includes           []string
	Events *Event
	Http   *Http
}

func (s *Configuration) ToString() string {
	result := ""
	// if s.User != "" {
	// 	result += fmt.Sprintf("user %v;\n", s.User)
	// }
	// if s.PID != "" {
	// 	result += fmt.Sprintf("pid %v;\n", s.PID)
	// }
	// if s.WorkerProcesses != "" {
	// 	result += fmt.Sprintf("worker_processes %v;\n", s.WorkerProcesses)
	// }
	// if s.ErrorLog != "" {
	// 	result += fmt.Sprintf("error_log %v;\n", s.ErrorLog)
	// }
	// if s.AccessLog != "" {
	// 	result += fmt.Sprintf("access_log %v;\n", s.AccessLog)
	// }
	// if s.WorkerRlimitNoFile != "" {
	// 	result += fmt.Sprintf("worker_rlimit_nofile %v;\n", s.WorkerRlimitNoFile)
	// }
	// if len(s.Includes) > 0 {
	// 	for _, include := range s.Includes {
	// 		result += fmt.Sprintf("include %v;\n", include)
	// 	}
	// }

	// if s.Events != nil {
	// 	result += s.Events.ToString()
	// }
	return result
}

func (s *Configuration) Parse(path string) error {
	// if !helper.FileExists(path) {
	// 	return errors.New("File " + path + " was not found.")
	// }

	// file, err := os.Open(path)
	// if err != nil {
	// 	return err
	// }

	// defer file.Close()

	// reader := bufio.NewReader(file)
	// var line string
	// for {
	// 	line, err = reader.ReadString('\n')
	// 	if err != nil && err != io.EOF {
	// 		break
	// 	}

	// 	if HasProperty(line, "user") {
	// 		s.User = GetPropertyValue(line, "user")
	// 		continue
	// 	}
	// 	if HasProperty(line, "worker_processes") {
	// 		s.WorkerProcesses = GetPropertyValue(line, "worker_processes")
	// 		continue
	// 	}
	// 	if HasProperty(line, "error_log") {
	// 		s.ErrorLog = GetPropertyValue(line, "error_log")
	// 		continue
	// 	}
	// 	if HasProperty(line, "access_log") {
	// 		s.AccessLog = GetPropertyValue(line, "access_log")
	// 		continue
	// 	}
	// 	if HasProperty(line, "pid") {
	// 		s.PID = GetPropertyValue(line, "pid")
	// 		continue
	// 	}
	// 	if HasProperty(line, "worker_rlimit_nofile") {
	// 		s.WorkerRlimitNoFile = GetPropertyValue(line, "worker_rlimit_nofile")
	// 		continue
	// 	}
	// 	if HasProperty(line, "include") {
	// 		if s.Includes == nil {
	// 			s.Includes = make([]string, 0)
	// 		}
	// 		include := GetPropertyValue(line, "include")
	// 		s.Includes = append(s.Includes, include)
	// 		continue
	// 	}

	// 	if err != nil {
	// 		break
	// 	}
	// }
	// if err != io.EOF {
	// 	return err
	// }

	return nil
}
