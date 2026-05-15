import React, { useState, useEffect, useMemo } from 'react';
import { 
  PlusCircle, 
  Save, 
  X, 
  Search, 
  Layers, 
  Edit2, 
  Trash2, 
  PackageOpen, 
  CheckCircle, 
  AlertCircle,
  ChevronsUpDown
} from 'lucide-react';
import './App.css';
import { Item } from './types';

const API_BASE = '/items';

function App() {
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState(true);
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState({ id: '', name: '', description: '' });
  const [searchTerm, setSearchTerm] = useState('');
  const [sortConfig, setSortConfig] = useState<{ key: keyof Item; direction: 'asc' | 'desc' }>({
    key: 'name',
    direction: 'asc'
  });
  const [toast, setToast] = useState<{ message: string; type: 'success' | 'error'; show: boolean }>({
    message: '',
    type: 'success',
    show: false
  });

  // Fetch items on mount
  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    setLoading(true);
    try {
      const response = await fetch(API_BASE);
      const data = await response.json();
      setItems(data || []);
    } catch (error) {
      showToast('Gagal memuat data', 'error');
    } finally {
      setLoading(false);
    }
  };

  const showToast = (message: string, type: 'success' | 'error' = 'success') => {
    setToast({ message, type, show: true });
    setTimeout(() => setToast(prev => ({ ...prev, show: false })), 3000);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const url = isEditing ? `${API_BASE}/${formData.id}` : API_BASE;
    const method = isEditing ? 'PUT' : 'POST';

    try {
      const response = await fetch(url, {
        method,
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          name: formData.name,
          description: formData.description
        })
      });

      if (response.ok) {
        showToast(isEditing ? 'Item berhasil diperbarui' : 'Item berhasil ditambahkan');
        resetForm();
        fetchItems();
      } else {
        throw new Error();
      }
    } catch (error) {
      showToast('Gagal menyimpan item', 'error');
    }
  };

  const handleEdit = (item: Item) => {
    setFormData(item);
    setIsEditing(true);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  const handleDelete = async (id: string) => {
    if (!confirm('Hapus item ini selamanya?')) return;
    try {
      const response = await fetch(`${API_BASE}/${id}`, { method: 'DELETE' });
      if (response.ok) {
        showToast('Item berhasil dihapus');
        fetchItems();
      } else {
        throw new Error();
      }
    } catch (error) {
      showToast('Gagal menghapus item', 'error');
    }
  };

  const resetForm = () => {
    setFormData({ id: '', name: '', description: '' });
    setIsEditing(false);
  };

  const handleSort = (key: keyof Item) => {
    setSortConfig(prev => ({
      key,
      direction: prev.key === key && prev.direction === 'asc' ? 'desc' : 'asc'
    }));
  };

  // Filter and Sort data
  const filteredAndSortedItems = useMemo(() => {
    let result = items.filter(item => 
      item.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      item.description.toLowerCase().includes(searchTerm.toLowerCase())
    );

    result.sort((a, b) => {
      const valA = a[sortConfig.key].toLowerCase();
      const valB = b[sortConfig.key].toLowerCase();
      if (valA < valB) return sortConfig.direction === 'asc' ? -1 : 1;
      if (valA > valB) return sortConfig.direction === 'asc' ? 1 : -1;
      return 0;
    });

    return result;
  }, [items, searchTerm, sortConfig]);

  return (
    <div className="container">
      <header>
        <h1>Task Master Pro <span style={{fontSize: '1rem', verticalAlign: 'middle', opacity: 0.7}}>React Edition</span></h1>
        <p>Kelola tugas dan inventaris Anda dengan elegan & reaktif</p>
      </header>

      <div className="main-grid">
        <div className="sidebar">
          <div className="card">
            <h2>
              {isEditing ? <Edit2 size={20} /> : <PlusCircle size={20} />}
              {isEditing ? 'Edit Item' : 'Tambah Item'}
            </h2>
            <form onSubmit={handleSubmit}>
              <div className="form-group">
                <label>Nama Item</label>
                <input 
                  type="text" 
                  value={formData.name}
                  onChange={e => setFormData({...formData, name: e.target.value})}
                  placeholder="Contoh: Belajar React TS" 
                  required 
                />
              </div>
              <div className="form-group">
                <label>Deskripsi</label>
                <textarea 
                  rows={4} 
                  value={formData.description}
                  onChange={e => setFormData({...formData, description: e.target.value})}
                  placeholder="Apa yang ingin Anda lakukan?" 
                  required 
                />
              </div>
              <button type="submit" className="btn btn-primary">
                <Save size={18} />
                <span>{isEditing ? 'Update Item' : 'Simpan Item'}</span>
              </button>
              {isEditing && (
                <button type="button" className="btn btn-secondary" onClick={resetForm}>
                  <X size={18} /> Batal
                </button>
              )}
            </form>
          </div>
        </div>

        <div className="content">
          <div className="card">
            <h2><Layers size={20} /> Data Explorer</h2>
            
            <div className="table-controls">
              <div className="search-wrapper">
                <Search size={18} className="search-icon" />
                <input 
                  type="text" 
                  placeholder="Cari berdasarkan nama atau deskripsi..." 
                  value={searchTerm}
                  onChange={e => setSearchTerm(e.target.value)}
                />
              </div>
              <div className="id-badge">{filteredAndSortedItems.length} Items</div>
            </div>

            {loading ? (
              <div className="loading-spinner"></div>
            ) : filteredAndSortedItems.length === 0 ? (
              <div className="empty-state">
                <PackageOpen size={48} style={{ marginBottom: '10px' }} />
                <p>Tidak ada data ditemukan.</p>
              </div>
            ) : (
              <div className="table-container">
                <table>
                  <thead>
                    <tr>
                      <th onClick={() => handleSort('name')}>
                        Nama <span className="sort-icon"><ChevronsUpDown size={14} /></span>
                      </th>
                      <th onClick={() => handleSort('description')}>
                        Deskripsi <span className="sort-icon"><ChevronsUpDown size={14} /></span>
                      </th>
                      <th>ID</th>
                      <th style={{ textAlign: 'right' }}>Aksi</th>
                    </tr>
                  </thead>
                  <tbody>
                    {filteredAndSortedItems.map(item => (
                      <tr key={item.id}>
                        <td><div className="item-name">{item.name}</div></td>
                        <td><div className="item-desc" title={item.description}>{item.description}</div></td>
                        <td><span className="id-badge">{item.id.substring(0, 8)}...</span></td>
                        <td>
                          <div className="item-actions">
                            <button className="btn-icon btn-edit" onClick={() => handleEdit(item)}>
                              <Edit2 size={18} />
                            </button>
                            <button className="btn-icon btn-delete" onClick={() => handleDelete(item.id)}>
                              <Trash2 size={18} />
                            </button>
                          </div>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )}
          </div>
        </div>
      </div>

      <div className={`toast ${toast.show ? 'show' : ''} ${toast.type === 'success' ? 'toast-success' : 'toast-error'}`}>
        {toast.type === 'success' ? <CheckCircle size={20} /> : <AlertCircle size={20} />}
        {toast.message}
      </div>
    </div>
  );
}

export default App;
