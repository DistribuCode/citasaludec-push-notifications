const mongoose = require('mongoose');

const doctorSchema = new mongoose.Schema({
  nombre: { type: String, required: true },
  especialidad: { type: String, required: true },
  telefono: { type: String },
  email: { type: String, unique: true },
  disponible: { type: Boolean, default: true }
}, {
  timestamps: true
});

module.exports = mongoose.model('Doctor', doctorSchema);
