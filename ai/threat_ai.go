// ai/threat_ai.go
type ThreatModel struct {
	Classifier *torch.Model
	Features   []string // размер, задержка, TTL, частота
}

func (tm *ThreatModel) UpdateWith(threat *ThreatReport) {
	tm.TrainOn(threat.Features, "attack")
	tm.AdaptParameters() // увеличить jitter, padding
}

func (tm *ThreatModel) PredictRisk() float64 {
	return tm.Classifier.Predict(currentTrafficPattern)
}