const Doctor = require('../models/Doctor');

// Crear médico
exports.createDoctor = async (req, res) => {
  try {
    const nuevoDoctor = new Doctor(req.body);
    const guardado = await nuevoDoctor.save();
    res.status(201).json(guardado);
  } catch (error) {
    res.status(400).json({ error: error.message });
  }
};

// Listar médicos
exports.getDoctors = async (req, res) => {
  try {
    const doctores = await Doctor.find();
    res.json(doctores);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};
